s = File.read("input.txt") |> elem(1)

s |> String.split("\n", trim: true) |> Enum.map(
  fn x -> 
    ["Card "<>n|[ cards ]]=String.split(x,":")
    [win,mine] = cards |> String.split("|",trim: true) |> Enum.map(fn x -> String.split(x," ",trim: true) |> Enum.map(&String.to_integer(&1) )end)
    mine |> Enum.filter(&(&1 in win)) |> length
      
end) |> Enum.map(&( if &1 >= 1, do: :math.pow(2,(&1-1)), else: 0)) |> Enum.sum()
