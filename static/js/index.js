const html = document.querySelector("html");
const nav = document.querySelector(".nav");
const hamburger = document.querySelector(".hamburger");
//////////////////////////////////////
// Mobile Navigation
//////////////////////////////////////
if (hamburger !== null) {
    hamburger.addEventListener("click", () => {
        nav.classList.toggle("open");
        hamburger.classList.toggle("close");
        html.classList.toggle("h-screen");
        html.classList.toggle("overflow-y-hidden");
    });
}
