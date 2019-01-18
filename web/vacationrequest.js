'use strict';

const textFields = document.querySelectorAll(".mdc-text-field");
for (let i = 0; i < textFields.length; i++) {
    mdc.textField.MDCTextField.attachTo(textFields[i]);
}

const buttons = document.querySelectorAll(".mdc-button");
for (let i = 0; i < buttons.length; i++) {
    mdc.ripple.MDCRipple.attachTo(buttons[i]);
}

const inputs = document.querySelectorAll("input, select, textarea");
const submitBtn = document.getElementById("submit");

let mode = document.body.dataset.mode;
updateUI();

function updateUI (){
    switch(mode) {
        case "new":
            mode = newMode;
            break;
        case "edit":
            // code block
            break;
        default:
            // show
            for (let i = 0; i < inputs.length; i++) {
                inputs[i].setAttribute("disabled","")
            }
    }
}

submitBtn.addEventListener("click",function (e) {
    console.log("Click");
    e.preventDefault();
})