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

def find_badge(group:)
  (group[0].split("") & group[1].split("") & group[2].split("")).join
end

sum_of_priorities = 0

input = File.open('input.txt').read

input.split("\n").each_slice(3) do |group|
  badge = find_badge(group: group)
  sum_of_priorities += alpha_to_int[badge]
end

puts sum_of_priorities

