#!/usr/bin/env ruby

input = File.open('input.txt').read

def parse_crate_line(line:, size:)
  ar = []
  size.times do |i|
    ar << line.slice((4*(i+1)) - 3)
  end
  ar
end

def parse_starting_stack(input:)
  starting_stack_hash = {}
  starting_stack = []

  input.each_line do |line|
    break if line == "\n"
    starting_stack << line
  end

  starting_stack_hash.tap do |h|
    starting_stack.last.strip.split.each do |s|
      h[s] = []
    end
  end
  *crates, last = starting_stack

  stacks = last.strip.split

  crates.each do |c|
    parse_crate_line(line: c.chop, size: stacks.size).each_with_index do |cc, i|
      starting_stack_hash[(i+1).to_s] = [cc, *starting_stack_hash[(i+1).to_s]] - [" ", nil]
    end
  end
  starting_stack_hash
end

def parse_moves(move:)
  mv = move.split
  [mv[1].to_i, mv[3], mv[5]]
end

passed_headers = false

starting_stack = parse_starting_stack(input: input)

input.each_line do |line|
  if line == "\n"
    passed_headers = true
    next
  end
  next unless passed_headers
  p = parse_moves(move: line)

  p p
  c = starting_stack[p[1]].pop(p[0])
  starting_stack[p[2]] = *starting_stack[p[2]], *c
end


answer = []
starting_stack.each do |k,v|
  answer << v.last
end

puts "The puzzle answer is #{answer.join}"
