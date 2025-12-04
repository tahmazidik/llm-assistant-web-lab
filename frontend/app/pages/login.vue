<script setup lang="ts">
import LoginForm from '~/widgets/auth/login-form'
definePageMeta({
  middleware: 'guest'
})
</script>

<template>
  <LoginForm />
</template>

<style scoped lang="scss">

</style>