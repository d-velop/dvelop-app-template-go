'use strict';

// component instantiation
mdc.topAppBar.MDCTopAppBar.attachTo(document.querySelector(".mdc-top-app-bar"));
const list = mdc.list.MDCList.attachTo(document.querySelector(".mdc-list"));
const menue = mdc.menu.MDCMenu.attachTo(document.querySelector(".mdc-menu"));
const snackbar = mdc.snackbar.MDCSnackbar.attachTo(document.querySelector(".mdc-snackbar"));
let selectedItem = document.querySelector(".mdc-list-item")

// event listeners
function handleMenuClick(state, icon) {
    const r = new XMLHttpRequest();
    r.addEventListener("load", function () {
        if (r.status == 200 || r.status == 204) {
            selectedItem.parentElement.dataset.state = state;
            selectedItem.querySelector(".state-icon").innerHTML = icon;
        } else {
            snackbar.labelText = "Request failed. Server returned " + r.status + ". Please try again in 5 seconds.";
            snackbar.open();
        }
    });
    r.addEventListener("error", function () {
        snackbar.labelText = "Request failed. Please try again in 5 seconds.";
        snackbar.open();
    });
    r.open("PATCH", selectedItem.href);
    r.setRequestHeader("Content-Type", "application/merge-patch+json");
    r.send(JSON.stringify({
        state: state
    }));
}

document.getElementById("menu_confirm").addEventListener("click", function(){
    handleMenuClick("confirmed", "check_circle");
});

document.getElementById("menu_reject").addEventListener("click", function(){
    handleMenuClick("rejected", "block");
});

document.getElementById("menu_cancel").addEventListener("click", function (){
    handleMenuClick("cancelled", "cancel");
});

const moreIconBtns = document.querySelectorAll(".mdc-icon-button");
for (let i = 0; i < moreIconBtns.length; i++) {
    const elmRipple = mdc.ripple.MDCRipple.attachTo(moreIconBtns[i]);
    elmRipple.unbounded = true;

    moreIconBtns[i].addEventListener("click", function (e) {
        e.stopPropagation();
        menue.open = !menue.open;
        selectedItem = e.currentTarget.parentElement.getElementsByTagName("a")[0];
        const btnElementRect = e.currentTarget.getBoundingClientRect();
        menue.setAbsolutePosition(btnElementRect.left, btnElementRect.top);
    });
}

