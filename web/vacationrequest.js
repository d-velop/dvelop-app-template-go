'use strict';

var mTextFields = [].map.call(document.querySelectorAll('.mdc-text-field'), function (el) {
    return mdc.textField.MDCTextField.attachTo(el);
});

var mButtons = [].map.call(document.querySelectorAll('.mdc-button'), function (el) {
    return mdc.ripple.MDCRipple.attachTo(el);
});

var mSelect = mdc.select.MDCSelect.attachTo(document.querySelector('#cat-select'));

var fulltextSearch = document.querySelector('#fulltext__enter-search-term');
var showResults = document.querySelector('#show-results');
var clear = document.querySelector('#clear');
var categorySelect = document.querySelector('#cat-select');

showResults.addEventListener('click', function (e) {
    e.preventDefault();
    console.log('Here are your results.');
    window.location = "results.html";
});

clear.addEventListener('click', function (e) {
    mTextFields.forEach(function (element) {
        element.value = '';
    });

    mSelect.selectedIndex = 0;
});

fulltextSearch.addEventListener('keyup', function (e) {
    e.preventDefault();

    if (e.keyCode === 13) {
        console.log('Here are your results.');
    }
});