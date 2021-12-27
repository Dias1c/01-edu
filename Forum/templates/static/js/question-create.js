let TabRadioButtons = document.getElementsByName("tab");


TabRadioButtons.forEach((rb)=>{
    const querySelector = rb.getAttribute("data-show-block");
    const tabBlock = document.querySelector(querySelector);

    // TODO CHECK EVENTS
    // rb.addEventListener('change', (e) => {
    //     if (rb.checked) {
    //         tabBlock.style.display = "block";
    // console.log("Block");        
    //     } else {
    //         tabBlock.style.display = "none";
    // console.log("None");
    //     } 
    // })
    // rb.addEventListener('deselect', () => {
    //     // if (rb.checked) {
    //         tabBlock.style.display = "none";
    //     // } 
    //     console.log("none");
    // })
    console.log(tabBlock);
});