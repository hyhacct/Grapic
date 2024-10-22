<template>
    <div>

        <div v-if="store.loader_content">
            <v-skeleton-loader type="paragraph@7"></v-skeleton-loader>
        </div>

        <div v-else>
            <highlightjs autodetect :code="store.respond_content">
            </highlightjs>
        </div>
    </div>
</template>


<script setup>
import { useStoreBranchRead } from "@/store/store_branch_read"

const store = useStoreBranchRead()

</script>