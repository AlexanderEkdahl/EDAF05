require 'rgl/adjacency'
require 'rgl/dot'
require 'rgl/traversal'

def distance(graph, a, b)
  distance = 0

  graph.bfs_iterator(a).each do |vertex|
    if vertex == b
      return distance
    end
    distance += 1
  end

  return -1
end

incoming = {}
graph = RGL::DirectedAdjacencyGraph.new

File.foreach(ARGV[0]).each do |line|
  word = line.chomp

  graph.add_vertex(word)

  wordx2 = word * 2
  incoming_keys = (0..word.length).map { |x| wordx2[x..x+3].chars.sort.join }

  incoming_keys.each do |key|
    if incoming.has_key?(key)
      incoming[key] << word
    else
      incoming[key] = [word]
    end
  end
end

graph.each do |word|
  outgoing_key = word[1..-1].chars.sort.join

  graph.add_vertex(word)

  incoming[outgoing_key].each do |other|
    if other != word
      graph.add_edge(word, other)
    end
  end
end

while input = STDIN.gets
  puts distance(graph, *input.chomp.split)
end

graph.write_to_graphic_file('svg')
