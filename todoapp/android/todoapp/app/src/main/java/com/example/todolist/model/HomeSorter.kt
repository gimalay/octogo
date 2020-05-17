package com.example.todolist.model

import com.example.todolist.model.ViewModelOuterClass.ViewModel

data class HomeSorter(
    val field: Int,
    val isDesc: Boolean
) {
    companion object {
        val default = HomeSorter(
            field = ViewModel.Home.Project.ID_FIELD_NUMBER,
            isDesc = true
        )

        private val projectFields = listOf(
            ViewModel.Home.Project.ID_FIELD_NUMBER,
            ViewModel.Home.Project.NAME_FIELD_NUMBER
        )

        val projectFieldTitles = listOf("ID", "Name")

        val fieldByTitle: Map<String, Int> = (projectFieldTitles zip projectFields).toMap()
    }
}