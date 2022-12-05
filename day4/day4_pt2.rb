#!/usr/bin/env ruby

count = 0
input = File.open('input.txt').read

input.each_line do |line|
  pairs = line.split(",")
  first_elf = eval(pairs.first.split("-").join("..")).to_a
  second_elf = eval(pairs.last.split("-").join("..")).to_a
  count +=1 unless (first_elf & second_elf).empty?
end

puts count
