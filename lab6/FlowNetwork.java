import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;

public class FlowNetwork {
	private HashMap<Integer, ArrayList<Edge>> adjacencyMatrix;
	private HashMap<Edge, Integer> flow;

	enum VertexState {
		White, Gray, Black
	}

	public static void main(String[] args) {
		FlowNetwork g = new FlowNetwork();
		g.parse();
		g.maxFlow(0, 54);
	}

	private void parse() {
		BufferedReader bf = null;
		int nbr;
		try {
			bf = new BufferedReader(new FileReader(new File("fixtures/rail.txt")));
			String line = bf.readLine();
			nbr = Integer.parseInt(line);
			for (int i = 0; i < nbr; i++) {
				line = bf.readLine();
				addVertex(i);
			}
			line = bf.readLine();
			nbr = Integer.parseInt(line);
			for (int i = 0; i < nbr; i++) {
				line = bf.readLine();
				String[] s = line.split(" ");
				if (Integer.parseInt(s[2]) == -1) {
					s[2] = String.valueOf(Integer.MAX_VALUE);
					// System.out.println("edge from" + Integer.parseInt(s[0]) +
					// " to " + Integer.parseInt(s[1]) + "is infifnittyty");
				}
				addEdge(Integer.parseInt(s[0]), Integer.parseInt(s[1]),
						Integer.parseInt(s[2]));
			}
		} catch (IOException e) {
			System.exit(0);
		}
//		 System.out.println(adj + " with size: " + adj.size());
	}

	public FlowNetwork() {
		adjacencyMatrix = new HashMap<Integer, ArrayList<Edge>>();
		flow = new HashMap<Edge, Integer>();
	}

	public void addVertex(int v) {
		adjacencyMatrix.put(v, new ArrayList<Edge>());
	}

	public ArrayList<Edge> getEdges(int s) {
		return adjacencyMatrix.get(s);
	}

	public void addEdge(int u, int v, int w) {
		Edge edge = new Edge(u, v, w);
		Edge redge = new Edge(v, u, w);
		edge.redge = redge;
		redge.redge = edge;
		adjacencyMatrix.get(u).add(edge);
		adjacencyMatrix.get(v).add(redge);
		flow.put(edge, 0);
		flow.put(redge, 0);
	}

	public ArrayList<Edge> findPath(int source, int sink, ArrayList<Edge> path,
			VertexState[] state) {
		state[source] = VertexState.Gray;
		if (source == sink) {
			return path;
		}
		for (Edge edge : getEdges(source)) {
			int residual = edge.capacity - flow.get(edge);
			if (residual > 0 && state[edge.sink] == VertexState.White) {
				@SuppressWarnings("unchecked")
				ArrayList<Edge> n = (ArrayList<Edge>) path.clone();
				n.add(edge);
				ArrayList<Edge> result = findPath(edge.sink, sink, n, state);
				if (result != null) {
					return result;
				}
			}
		}
		state[source] = VertexState.Black;

		return null;
	}

	public ArrayList<Edge> findPath(int source, int sink) {
		VertexState state[] = new VertexState[adjacencyMatrix.size()];
		for (int i = 0; i < adjacencyMatrix.size(); i++){
			state[i] = VertexState.White;
		}
		return findPath(source, sink, new ArrayList<Edge>(), state);
	}

	public void maxFlow(int source, int sink) {
		ArrayList<Edge> path = findPath(source, sink);
		while (path != null) {
			int flow = Integer.MAX_VALUE;
			for (Edge edge : path) {
				int x = edge.capacity - this.flow.get(edge);
				if (x < flow) {
					flow = x;
				}
			}
			for (Edge edge : path) {
				this.flow.put(edge, this.flow.get(edge) + flow);
				this.flow.put(edge.redge, this.flow.get(edge.redge) - flow);
			}
			path = findPath(source, sink);
		}

		HashSet<Edge> visited = new HashSet<Edge>();
		buildMinimumCut(source, visited);
		printMinimumCut(visited);

		int sum = 0;
		for (Edge edge : getEdges(source)) {
			sum += flow.get(edge);
		}

		System.out.println("Maximum flow: " + sum);
	}

	private void printMinimumCut(HashSet<Edge> visited) {
		HashSet<Integer> vertices = new HashSet<Integer>();
		for (Edge edge : visited) {
			vertices.add(edge.source);
		}

		for (Edge edge : visited) {
			if(!vertices.contains(edge.sink)) {
				System.out.println(edge);
			}
		}
	}

	public void buildMinimumCut(int source, HashSet<Edge> visited) {
		for (Edge edge : getEdges(source)) {
			if(visited.contains(edge)){
				continue;
			}
			visited.add(edge);
			if(flow.get(edge) != edge.capacity) {
				buildMinimumCut(edge.sink, visited);
			}
		}
	}

	public class Edge {
		private int source, sink;
		private int capacity;
		private Edge redge;

		public Edge(int source, int sink, int capacity) {
			this.source = source;
			this.sink = sink;
			this.capacity = capacity;
		}

		public String toString() {
			return String.valueOf(source) + "->" + String.valueOf(sink) + ":"
					+ String.valueOf(capacity);
		}
	}
}
