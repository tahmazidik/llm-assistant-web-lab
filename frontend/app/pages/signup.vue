<script setup lang="ts">
import { SignUpForm} from "~/widgets/auth/signup-form";

definePageMeta ({
  middleware: ['guest']
})
</script>

<template>
  <SignUpForm />
</template>

<style scoped lang="scss">

</style>