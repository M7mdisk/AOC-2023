res, s= File.read("input.txt")
ar = String.split(s,"\n")

ar |> List.delete_at(-1) |> Enum.map(fn x-> x|> String.to_charlist() |> Enum.filter(fn x-> x <65 end) |>(&[Enum.at(&1, 0), Enum.at(&1, -1)]).() |> List.to_integer() end) |> Enum.reduce( 0, &+/2)
