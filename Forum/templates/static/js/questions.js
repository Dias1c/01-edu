// Widgets like Search

// HEADER SEARCH
var cb_QuestionsFilter = document.getElementById("cb_QuestionsFilter");
var b_QuestionsFilter = document.getElementById("b_QuestionsFilter");


// cb_HeaderSearch_change - Shows or hide element
function cb_QuestionsFilter_change(e) {
    if (e.target.checked) {
        b_QuestionsFilter.classList.add("dflex");
    } else {
        b_QuestionsFilter.classList.remove("dflex");
        console.log("Haha")
    }
}
cb_QuestionsFilter.addEventListener("change", cb_QuestionsFilter_change);
