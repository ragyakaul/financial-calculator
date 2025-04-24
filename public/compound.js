//Get form data
let myChart
document.getElementById("sendToServer").addEventListener("click", async function(event){
    event.preventDefault();


    const formData = {
        initialPortfolioValue: parseFloat(document.getElementById("initialPortfolioValue").value),
        annualBaseContribution: parseFloat(document.getElementById("annualBaseContribution").value),
        inflationRate: parseFloat(document.getElementById("inflationRate").value),
        growthRate: parseFloat(document.getElementById("growthRate").value),
        yearsCompounding: parseInt(document.getElementById("yearsCompounding").value),
        targetFIREValue: parseFloat(document.getElementById("targetFIREValue").value),
        annualWithdrawalPercentage: parseFloat(document.getElementById("annualWithdrawalPercentage").value),
    }

    // Send data to the server using Fetch API

    let response = await fetch("http://localhost:8000/api/v1/compounding", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    });

    let serverResponse = await response.json()

    if (myChart){
        myChart.destroy()
    }
    const phasePlugin = {
        id: 'phaseBackground',
        beforeDraw: (chart) => {
            const {ctx, chartArea: {top, bottom, left, right}, scales: {x}} = chart;

            const targetYear = serverResponse.findIndex(v => v >= formData.targetFIREValue);
            const targetX = x.getPixelForValue(targetYear);

            // Contributing phase (left side)
            ctx.fillStyle = 'rgb(175, 73, 73)'
            ctx.fillRect(left, top, targetX - left, bottom - top);

            // Withdrawing phase (right side)
            ctx.fillStyle = 'rgb(105, 204, 125)'
            ctx.fillRect(targetX, top, right - targetX, bottom - top);

            // Add text
            ctx.fillStyle = 'white';
            ctx.font = 'bold 14px sans-serif';
            ctx.fillText('Working', left + 10, top + 20);
            ctx.fillText('Retired', targetX + 50, top + 20);
        }
    }

    const ctx = document.getElementById('myChart').getContext('2d');
    myChart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: Array.from({ length: 75 }, (_, i) => i + 1),
            datasets: [{
                label: 'Portfolio Value',
                data: serverResponse,
                borderColor: 'white',
                borderWidth: 2, 
                fill: false
            }]
        },
        options: {
            responsive: true,
            scales: {
                x: {title: {display: true, text: 'Year '}},
                y: {title: {display: true, text: 'Value ($)'}}
            }, 
            plugins: {
                legend: {display:true}
            }
        },
        plugins: [phasePlugin]
    })
    console.log(serverResponse)
});

