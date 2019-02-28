fetch("/api/vouchers").then(e => {
    return e.json()
}).then(data => {
    console.log(data);
    let i = 1;
    data.forEach(voucher => {
        i++;
        document.querySelector("#app").innerHTML += "<div class=\"voucher " + (i % 2 === 0) + "\"><img class=\"voucher-icon\" src=\"/assets/icon.png\">" + voucher.Code + "</div>";
    });
});

document.querySelector("#print").addEventListener('click', () => window.print());