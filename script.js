async function fetchAsync (URL) {
  let response = await fetch(URL);
  let data = await response.json();
  return data;
}

document.getElementsByClassName("submit")[0].addEventListener('click', function (){
    console.log(fetchAsync("http://localhost:8080/lastid"))
})
