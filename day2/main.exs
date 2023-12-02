my_function = fn(x) ->
  ["Game "<>n|[rest]]=String.split(x,":")
  rest = rest |> String.split(";") |> Enum.map(&String.split(&1,",")) |> List.flatten() |> Enum.map(&String.trim/1) 
  mp = Enum.reduce(rest, %{red: 0, green: 0, blue: 0}, fn c, mx ->   
        [count, color] = String.split(c) 
        Map.update(mx, String.to_atom(color), count, &max(&1, String.to_integer(count)))
      end)
  if mp.red <= 12 and mp.green <= 13 and mp.blue <= 14, do: String.to_integer(n), else: 0
end
IO.puts(File.read("input.txt") |> elem(1) |> String.split("\n",trim: true) |> Enum.map(my_function) |> Enum.reduce( 0, &+/2))
