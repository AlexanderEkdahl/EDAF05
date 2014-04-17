require 'rgl/adjacency'
# require 'rgl/dot'
require 'rgl/traversal'

def distance(graph, a, b)
  # v = graph.instance_variable_get(:@vertice_dict)[a]
  iterator = graph.bfs_iterator(a)
  iterator.attach_distance_map

  iterator.each do |vertex|
    if vertex == b
      return iterator.distance_to_root(vertex) - 1
    end
  end

  return -1
end

incoming = {}
graph = RGL::DirectedAdjacencyGraph.new

File.foreach(ARGV[0]).each do |line|
  word = line.chomp

  graph.add_vertex(word)

  wordx2 = word * 2
  incoming_keys = (0..word.length).map { |x| wordx2[x..x + 3].chars.sort.join }

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

# graph.write_to_graphic_file('svg')
