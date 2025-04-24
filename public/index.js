


// Get form data

document.getElementById("sendToServer").addEventListener("click", async function(event) {
event.preventDefault();


const formData = {
    loan: parseFloat(document.getElementById("loanOnHouse").value),
    annualInterestRate: parseFloat(document.getElementById("currentAnnualHomeInterestRate").value),
    yearsLeftOnLoan: parseFloat(document.getElementById("yearsLeftOnLoan").value)
}

// Send data to the server using Fetch API

const response = await fetch("http://localhost:8000/api/v1/mortgage", {
    method: "POST",
    headers: {
        "Content-Type": "application/json"
    },
    body: JSON.stringify(formData)
});

const serverResponse = await response.json() 
document.getElementById("monthlyPrincipal").value = serverResponse.monthlyPrincipal 
document.getElementById("yearlyPrincipal").value = serverResponse.yearlyPrincipal
document.getElementById("monthlyInterest").value = serverResponse.monthlyInterest
document.getElementById("yearlyInterest").value = serverResponse.yearlyInterest
});
