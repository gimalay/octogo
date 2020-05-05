package com.example.todolist.ui.home

import androidx.compose.Composable
import androidx.ui.foundation.VerticalScroller
import androidx.ui.layout.Column
import androidx.ui.tooling.preview.Preview
import com.example.todolist.model.ViewModelOuterClass
import com.example.todolist.utils.newUUIDasByteString

@Composable
fun HomeBody(
    children: @Composable() () -> Unit
) {
    VerticalScroller {
        Column {
            children()
        }
    }
}

@Preview
@Composable
fun PreviewHomeActivityBody() {
    val projects = listOf(
        ViewModelOuterClass.ViewModel.Home.Project
            .newBuilder()
            .setID(newUUIDasByteString())
            .setName("Test project 1")
            .build(),
        ViewModelOuterClass.ViewModel.Home.Project
            .newBuilder()
            .setID(newUUIDasByteString())
            .setName("Test project 2")
            .build()
    )
    HomeBody{
        projects.forEach {
            Activity(
                text = it.name,
                onCopy = {},
                onRemove = {}
            )
        }
    }
}
