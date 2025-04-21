module RuleMap = struct
  let spacing = [
    ("mt-4", "margin-top: 1rem;");
    ("py-12", "padding-top: 3rem; padding-bottom: 3rem;");
    ("px-4", "padding-left: 1rem; padding-right: 1rem;");
    ("pt-2", "padding-top: 0.5rem;");
  ]

  let text = [
    ("text-center", "text-align: center;");
    ("font-bold", "font-weight: 700;");
  ]

  let colors = [
    ("bg-blue-500", "background-color: #3b82f6;");
    ("text-white", "color: #ffffff;");
  ]

  let all = spacing @ text @ colors
end

module Variants = struct
  let handlers = [
    ("hover", (fun cls rule ->
      Printf.sprintf ".hover\\:%s:hover { %s }" cls rule));

    ("sm", (fun cls rule ->
      Printf.sprintf "@media (min-width: 640px) { .sm\\:%s { %s } }" cls rule));
  ]

  let apply variant cls rule =
    match List.assoc_opt variant handlers with
    | Some f -> f cls rule
    | None -> ""
end

let lookup_rule base =
  match List.assoc_opt base RuleMap.all with
  | Some rule -> Printf.sprintf ".%s { %s }" base rule
  | None -> ""

let class_to_rule full_class =
  match String.split_on_char ':' full_class with
  | [base] -> lookup_rule base
  | [variant; base] ->
      (match List.assoc_opt base RuleMap.all with
      | Some rule -> Variants.apply variant base rule
      | None -> "")
  | _ -> ""
