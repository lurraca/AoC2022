# AoC2022 Day1 - part 2

# Object because why not
class Elf
  attr_accessor :food_items
  def initialize(food_items:)
    @food_items = food_items
  end

  def total_calories
    food_items.sum
  end
end

# Hold all Elf classes
elves = []

# Used to store the items of a single Elf at time
items = []

input = File.open('input.txt').read

input.each_line do |line|
  if line == "\n" # Next line will be the calories of a new elf
    elves << Elf.new(food_items: items)
    items = [] # Clean transient variable for next Elf
  else
    items << line.to_i
  end
end

puts "The sum of calories of the top 3 elves is: #{elves.map(&:total_calories).max(3).sum}"

# Given my input, the sum of the calories that the top 3 elves are holding is: 203002
