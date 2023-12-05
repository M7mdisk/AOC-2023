s = File.read("input.txt") |> elem(1)

[seeds | maps] = String.split(s, "\n\n", trim: true)
"seeds:" <> seeds = seeds
seeds = seeds |> String.split(" ", trim: true) |> Enum.map(&String.to_integer(&1))

maps =
  Enum.map(maps, fn x ->
    String.split(x, "\n")
    |> Enum.slice(1..-1)
    |> Enum.map(fn x -> String.split(x, " ") |> Enum.map(&String.to_integer(&1)) end)
  end)

res =
  Enum.map(seeds, fn sd ->
    Enum.reduce(maps, sd, fn map, seed ->
      Enum.reduce(map, seed, fn [dest, src, len], acc ->
        if acc in src..(src + len) and acc == seed do
          dest + acc - src
        else
          acc
        end
      end)
    end)
  end)
  |> Enum.min()

IO.puts(res)
