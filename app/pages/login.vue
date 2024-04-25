<script setup>

    const logIn = () => {
        console.log('logged in');
        navigateTo('/');
    }
</script>

<template>
    <div class='page'>
        <form class='form'>
            <input placeholder='Username'></input>
            <input placeholder='Password' type='password'></input>
            <button class='btn' type=submit @click='logIn'>Log in</button>
        </form>
    </div>
</template>

<style scoped>
.title {
    margin: 16px;
}

.form {
    display: flex;
    flex-direction: column;
    gap: 4px;
}

input {
    padding: 4px;
    min-width: 20%;
}

.btn {
    padding: 4px;
    align-self: center;
    width: 60%;
    margin-top: 8px;
}

</style>
