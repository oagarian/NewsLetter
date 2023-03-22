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

function validateEmail(email) {
  const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return re.test(email);
}

function sendData(email) {
  url = "http://localhost:8080/register"
  requestGET("http://localhost:8080/lastid").then((value) => {
    var atualValue = parseInt(value) + 1;
    const data = { ID: atualValue, Email: email };

    fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    })
      .then(response => response.json())
      .then(json => console.log(json))
      .catch(error => console.error(error));
  })
  
}

document.getElementsByClassName("submit")[0].addEventListener('click', function(){
  inputValue = document.getElementById("email").value;
  if (validateEmail(inputValue)) {
    sendData(inputValue)
    window.alert("Sucessfully submitted")
  } else {
    window.alert("Invalid email")
  }
})

