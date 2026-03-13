function jumping(bool) {
  const sp = document.getElementById("jumping_indicator");
  if (bool == true) {
    sp.classList.add("jumping");
    return;
  }

  sp.classList.remove("jumping");
}

function spin(bool) {
  const sp = document.getElementById("spinning_indicator");
  if (bool == true) {
    sp.classList.add("spinner_round");
    return;
  }

  sp.classList.remove("spinner_round");
}


