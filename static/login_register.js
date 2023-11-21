var registerBtn = document.getElementById("reg_btn");
var loginBtn = document.getElementById("login_btn");

// Registration
registerBtn.addEventListener('click', function() {
    var username = document.getElementById("reg_username").value;
    var password = document.getElementById("reg_password").value;
    
    $.ajax({
        url: "/register",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({"username": username, "password": password}),
        success: function(data) {
            console.log("registration successful!", data);
        },
        error: function(jqXHR, textStatus, errorThrown) {
            console.log("registration error", textStatus, errorThrown); 
        },
    });
});



