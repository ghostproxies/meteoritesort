// Copyright William Stafford Parsons

func ImpactSort(elements []int) {
  var element int
  gap := len(elements)
  var i int
  var j int

  for gap > 15 {
    gap = ((gap >> 2) - (gap >> 4)) + (gap >> 3)
    i = gap

    for i < len(elements) {
      element = elements[i]
      j = i

      for j >= gap && elements[j - gap] > element {
        elements[j] = elements[j - gap]
        j -= gap
      }

      elements[j] = element
      i++
    }
  }

  i = 1
  gap = 0

  for i < len(elements) {
    element = elements[i]
    j = i

    for j > 0 && elements[gap] > element {
      elements[j] = elements[gap]
      j = gap
      gap--
    }

    elements[j] = element
    gap = i
    i++
  }
}
