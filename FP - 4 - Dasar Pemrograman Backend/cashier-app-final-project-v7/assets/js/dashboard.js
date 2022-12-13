$(document).ready(function() {
    var isFormUpdateImage = false
    $(".btn-update").click(function() {
        var self = $(this),
            btnImage = self.parent("div.btn-update-image"),
            formUpdateImage = btnImage.prev();
        formUpdateImage.toggle("active:block active:mx-2")
        if (!isFormUpdateImage) {
            isFormUpdateImage = true
            self.html("Cancel")
        } else {
            isFormUpdateImage = false
            self.html("Update Image")
        }
    })
})

function displayQty(listData, checkId, qtyId) {
    var checkBox = document.getElementById(checkId);
    console.log(checkBox.checked)
    var qty = document.getElementById(qtyId);
    if (checkBox.checked == true) {
        qty.style.display = "block";
        checkBox.value = listData + "," + qty.value
    } else {
        qty.style.display = "none";
    }
}