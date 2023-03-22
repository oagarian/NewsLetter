async function requestGET(URL) {
  const response = await fetch(URL);
  const data = await response.text();
  const value = Object.values(data)[0];
  return value
}

const counter = document.getElementsByClassName("counter")[0]
requestGET("http://localhost:8080/lastid").then((value) => {
  counter.textContent = value;
})


