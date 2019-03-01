fetch("/api/vouchers").then(e => {
    return e.json()
}).then(data => {
    console.log(data);
    let i = 1;
    data.forEach(voucher => {
        i++;
        let n = "";
        if (i % 24 === 0 || i % 25 === 0) {
            n = "page";
        }
        document.querySelector("#app").innerHTML += "<div class=\"voucher " + (i % 2 === 0) + " " + n + "\"><img class=\"voucher-icon\" src=\"/assets/icon.png\">" + voucher.Code + "</div>";
    });
});

document.querySelector("#print").addEventListener('click', () => window.print());