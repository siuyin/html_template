function spin_round(bool) {
  const sp = document.getElementById("jumping_indicator");
  if (bool == true) {
    sp.classList.add("spinner_round");
    return;
  }

  sp.classList.remove("spinner_round");
}
