s = """
Time:        53897698
Distance:   313109012141201
"""

[time | [dist]] =
  String.split(s, "\n", trim: true)
  |> Enum.map(&String.split(&1, " ", trim: true))
  |> Enum.map(fn x -> Enum.slice(x, 1..-1) |> Enum.map(&String.to_integer(&1)) end)

Enum.zip(time, dist)
  |> Enum.map(fn {t, d} ->
    Enum.reduce(0..t, fn hold, acc ->
    if hold * (t - hold) > d, do: acc + 1, else: acc
  end)
end)
