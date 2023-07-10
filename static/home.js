// Make 10 boxes 
// Grab box element

let boxes = {"row": 0, "col": 0}
function createBoxes() {
    let obj = document.body
    let FirstDiv = document.getElementById("first-div");
    if (!FirstDiv) {
        FirstDiv = document.createElement("div");
        FirstDiv.id = "first-div";
    } else {
        FirstDiv.innerHTML = "";
    }



    for (let j=0; j<boxes["row"]; j++) 
        {
            let SecondDiv = document.createElement("div");
            for (let i=0; i<boxes["col"]; i++) {
                // Create new element
                let newElement = document.createElement("div");
                // newElement.addEventListener("mouseover", handleHover);
                // Set class
                newElement.className = "box";
                newElement.id = j + "," + i;
                // Append to parent
                SecondDiv.appendChild(newElement);
            }
            FirstDiv.appendChild(SecondDiv);
        }
    obj.appendChild(FirstDiv);
}


function updateValue(value, id) {
    document.getElementById(id).textContent = value;
    boxes[id] = value;
    createBoxes();
}

const handleSubmit = (event) => {
    event.preventDefault();
    let row = boxes["row"];
    let col = boxes["col"];
    fetch('/api/init-matrix', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({row: row, col: col}),
    })
    .catch(err => {
        console.log(err);
    })
}   

const handleHover = (event) => {
    let objects = document.getElementsByClassName("box");
    objects.forEach(element => {
        element.style.backgroundColor = "red";
})
}