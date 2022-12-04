#!/usr/bin/env ruby

alphabet_lower_case = 'a'..'z'
alphabet_upper_case = 'A'..'Z'

alpha_to_int = {}


alphabet_lower_case.each do |letter|
  alpha_to_int[letter] = letter.ord - 96
end

alphabet_upper_case.each do |letter|
  alpha_to_int[letter] = letter.ord - 38
end

def split_rucksack(rucksack:)
  rucksack.partition(/.{#{rucksack.size/2}}/)[1,2]
end

sum_of_priorities = 0

input = File.open('input.txt').read

input.each_line do |line|
  first, second = split_rucksack(rucksack:line)
  intersection = first.split("") & second.split("")
  p intersection.join
  sum_of_priorities += alpha_to_int[intersection.join]
end

puts sum_of_priorities

