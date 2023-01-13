let LightDarkMode = document.querySelector(".LightDarkMode")
let body = document.querySelector("body")
LightDarkMode.addEventListener("click", () => {
    LightDarkMode.classList.toggle("active")
    body.classList.toggle("active")
})

// Password validation
// let password = document.getElementById("password")
// let confirm_password = document.getElementById("confirm_password");

// function validatePassword() {
//     if (password.value != confirm_password.value) {
//         confirm_password.setCustomValidity("Passwords Don't Match");
//     } else {
//         confirm_password.setCustomValidity('');
//     }
// }

// password.onchange = validatePassword;
// confirm_password.onkeyup = validatePassword;