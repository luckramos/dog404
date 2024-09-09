import './style.scss';

const toggleForm = document.querySelector('.toggle-form-input');
toggleForm.addEventListener('change', () => {
    const formWrapper = document.querySelector('.form-wrapper');
    const toggleText = document.querySelector('.toggle-text');
    if(toggleForm.checked){
        formWrapper.classList.add('open');
        toggleText.innerHTML = 'X';
        return;
    }

    formWrapper.classList.remove('open');
    toggleText.innerHTML = 'Entrar em contato';

})

const emailForm = document.querySelector('.email-form');
emailForm.addEventListener('submit', (e) => {
    e.preventDefault();
    const formData = new FormData(emailForm);
    fetch('http://localhost:8080/send-email', {
        method: 'POST',
        body: formData
    })
    .then(res => {
        res.json().then(data => {
            document.querySelector('.form-wrapper').innerHTML +=`
                <div class="success-message">
                    <p>${data.message}</p>
                </div>
            `
        })
    })
})