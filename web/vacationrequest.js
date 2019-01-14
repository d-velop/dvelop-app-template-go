'use strict';

const textFields = document.querySelectorAll(".mdc-text-field");
for (let i = 0; i < textFields.length; i++) {
    mdc.textField.MDCTextField.attachTo(textFields[i]);
}

const buttons = document.querySelectorAll(".mdc-button");
for (let i = 0; i < buttons.length; i++) {
    mdc.ripple.MDCRipple.attachTo(buttons[i]);
}
