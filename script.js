async function fetchAsync (URL) {
  fetch(URL)
  .then((response) => response.json())
  .then((json) => console(json))
}

document.getElementsByClassName("submit")[0].addEventListener('click', function (){
    fetchAsync("http://localhost:8080/lastid")

})
