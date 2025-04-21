open Css

let generate_css classes =
  classes
  |> List.map class_to_rule
  |> List.filter (fun rule -> rule <> "")
  |> String.concat "\n"