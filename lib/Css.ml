let class_to_rule = function
  | "text-center" -> ".text-center { text-align: center; }"
  | "font-bold" -> ".font-bold { font-weight: 700; }"
  | "mt-4" -> ".mt-4 { margin-top: 1rem; }"
  | "py-12" -> ".py-12 { margin-top: 12rem; margin-bottom: 12rem; }"
  | "bg-blue-500" -> ".bg-blue-500 { background-color: #3b82f6; }"
  | _ -> ""