class Universe
  def initialize(path)
    @history = [[],[],[]]
    @moons_pos = {}
    @moons_vel = {}

    File.read(path).lines.each_with_index do |l,i|
      @moons_pos[i] = l.scan(/[-0-9]+/).map(&:to_i)
      @moons_vel[i] = [0,0,0]
    end

    @moons_initial_pos = @moons_pos.clone
    @moons_initial_vel = @moons_vel.clone
  end

  def simulate
    snapshot
    apply_gravity
    apply_velocity
  end

  def find_periods
    periods = []
    @history.each.with_index do |points, i|
      periods << points.each_index.select {|i| points[i] == points[0]}[1]
    end
    periods.inject(:lcm)
  end

  def calculate_potential_energy
    @moons_pos.values.map {|pos| pos.map(&:abs).sum}
  end

  def calculate_kinetic_energy
    @moons_vel.values.map {|vel| vel.map(&:abs).sum}
  end

  def calculate_total_energy
    calculate_potential_energy.zip(calculate_kinetic_energy).map {|p, k| p * k}.sum
  end

  private

  def snapshot
    @history.each.with_index do |_, i|
      @history[i] << [@moons_pos[0][i], @moons_pos[1][i], @moons_pos[2][i], @moons_pos[3][i],
                      @moons_vel[0][i], @moons_vel[1][i], @moons_vel[2][i], @moons_vel[3][i]]
    end
  end

  def apply_gravity
    @moons_pos.keys.combination(2).each do |first, second|
      changes = @moons_pos[first].zip(@moons_pos[second]).map{|a,b| a <=> b}
      @moons_vel[first] = [@moons_vel[first],changes].transpose.map{|v| v.reduce(:-)}
      @moons_vel[second] = [@moons_vel[second], changes].transpose.map{|v| v.reduce(:+)}
    end
  end

  def apply_velocity
    @moons_pos.each do |k,v|
      @moons_pos[k] = [v,@moons_vel[k]].transpose.map(&:sum)
    end
  end
end

# part 1
universe = Universe.new('day12.txt')
1000.times do
  universe.simulate
end

p universe.calculate_total_energy

universe = Universe.new('day12.txt')
300000.times do
  universe.simulate
end

p universe.find_periods
