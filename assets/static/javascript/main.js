document.addEventListener('htmx:beforeSwap', evt => {
    const status = evt.detail.xhr.status;

    if (status === 404 || status === 500) {
        evt.detail.shouldSwap = true;
        evt.detail.isError = false;
    }
});