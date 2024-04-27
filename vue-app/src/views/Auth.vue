<template>
    <div class='page'>
        <h1>Welcome to DJmocracy</h1>
        <h2>In order to use this app you must sign up or log in</h2>
        <RouterLink to='/login'>Log in</RouterLink>
        <RouterLink to='/register'>Sign up</RouterLink>
        <RouterView />
    </div>
</template>

<style scoped>
h1 {
    color: white;
}
</style>
