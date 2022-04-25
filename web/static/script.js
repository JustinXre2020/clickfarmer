window.addEventListener('load', function () {
    updateClickLabels();
})

function colorClicked(color) {
    console.log("color clicked", color);
    // hmm
    localStorage.setItem(color, localStorage.getItem(color) != null ? parseInt(localStorage.getItem(color)) + 1 : 1)
 
    var xhttp = new XMLHttpRequest();
    var data = parseInt(localStorage.getItem(color))
    if (color == 'red') {
        xhttp.open("PUT", "/api/clicks/red");
    }
    else if (color == 'green') {
        xhttp.open("PUT", "/api/clicks/green");
    }
    else {
        xhttp.open("PUT", "/api/clicks/blue");
    }
    xhttp.send(data);
}

function updateClickLabels() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var response = JSON.parse(this.responseText);
            document.getElementById("color-label-red").innerHTML = "red: " + response.redClicks
            document.getElementById("color-label-green").innerHTML = "green: " + response.greenClicks
            document.getElementById("color-label-blue").innerHTML = "blue: " + response.blueClicks
        }
    };
    xhttp.open("GET", "/api/clicks", true);
    xhttp.send();
}
