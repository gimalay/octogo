package com.example.todolist.model

data class HomeFilter(
    val projectName: String,
    val projectId: String
) {
    companion object {
        val default = HomeFilter(
            projectId = "",
            projectName = ""
        )
    }
}