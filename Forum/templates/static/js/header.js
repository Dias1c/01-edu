// Widgets like Search
var cb_Search = document.getElementById("cb_search");
var b_Search = document.getElementById("b_search");

// cb_Search_change - Shows or hide element
function cb_Search_change(e) {
    if (e.target.checked) {
        b_Search.classList.add("dflex-on-md");
        cb_Search.classList.add("search-btn--checked");
    } else {
        b_Search.classList.remove("dflex-on-md");
        cb_Search.classList.remove("search-btn--checked");
    }
}

// ADDING EVENTS
cb_Search.addEventListener("change", cb_Search_change);
