const html = document.querySelector("html");
const scheme = localStorage.getItem("scheme");

document.addEventListener("htmx:beforeSwap", evt => {
    const status = evt.detail.xhr.status;

    if (status === 404 || status === 500) {
        evt.detail.shouldSwap = true;
        evt.detail.isError = false;
    }
});

if (scheme === "auto") switchAuto();
if (scheme === "light") switchLight();
if (scheme === "dark") switchDark();

function switchAuto() {
    html.style.setProperty("color-scheme", "light dark");
    localStorage.removeItem("scheme");
}

function switchLight() {
    html.style.setProperty("color-scheme", "light");
    localStorage.setItem("scheme", "light");
}

function switchDark() {
    html.style.setProperty("color-scheme", "dark");
    localStorage.setItem("scheme", "dark");
}

function switchMode(mode) {
    if (mode === "auto") {
        switchAuto();
        return;
    }

    html.style.setProperty("color-scheme", mode)
    localStorage.setItem("scheme", mode);
}