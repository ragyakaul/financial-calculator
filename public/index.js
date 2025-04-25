


// Get form data
let myChart
document.getElementById("sendToServer").addEventListener("click", async function(event) {
event.preventDefault();


const formData = {
    loan: parseFloat(document.getElementById("loanOnHouse").value),
    mandatoryMonthlyPayment: parseFloat(document.getElementById("mandatoryMonthlyPayment").value),
    optionalMonthlyPayment: parseFloat(document.getElementById("optionalMonthlyPayment").value),
    annualInterestRate: parseFloat(document.getElementById("currentAnnualHomeInterestRate").value),
    yearsLeftOnLoan: parseInt(document.getElementById("yearsLeftOnLoan").value)
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
const months = Array.from({length: serverResponse.interestValues.length}, (_, i) => `Month ${i + 1}`);

if (myChart){
    myChart.destroy()
}

const ctx = document.getElementById("mortgageChart").getContext("2d");
myChart = new Chart(ctx, {
    type: "line",
    data: {
        labels: months,
        datasets: [
            {
                label: "Interest Paid",
                data: serverResponse.interestValues,
                borderColor: "rgba(255, 99, 132, 1)",
                backgroundColor: "rgba(255, 99, 132, 0.2)",
                tension: 0.1
            },
            {
                label: "Principal Paid",
                data: serverResponse.principalValues,
                borderColor: "rgba(54, 162, 235, 1)",
                backgroundColor: "rgba(54, 162, 235, 0.2)",
                tension: 0.1
            }
        ]
    },
    options : {
        responsive: true,
        plugins: {
            title: {
                display: true,
                text: "Mortgage Payment Breakdown"
            }
        },
        scales:{
            y: {
                beginAtZero: false
            }
        }
    }
})
console.log(serverResponse)
});
