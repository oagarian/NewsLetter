submitButton = document.getElementsByClassName("submit")[0];
emailInput = document.getElementById("email")
subsCounter = document.getElementsByClassName("counter")[0];

submitButton.addEventListener("click", function(event) { 
    console.log(emailInput.value);
})