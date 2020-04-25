package com.example.todolist.ui

import androidx.compose.*
import androidx.ui.foundation.VerticalScroller
import androidx.ui.layout.Column
import androidx.ui.material.MaterialTheme
import androidx.ui.tooling.preview.Preview
import com.example.todolist.data.newUUIDasByteString
import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.google.protobuf.ByteString

@Composable
fun HomeActivity(
    projects: List<ViewModel.Home.Project>,
    onCopyProject: (name: String) -> Unit = {},
    onRemoveProject: (id: ByteString) -> Unit = {}
) {
    MaterialTheme(
        colors = lightThemeColors,
        typography = themeTypography
    ) {
        VerticalScroller {
            Column {
                projects.forEach {
                    Activity(
                        text = it.name,
                        onCopy = { onCopyProject(it.name) },
                        onRemove = { onRemoveProject(it.id) }
                    )
                }
            }
        }
    }
}

@Preview
@Composable
fun PreviewApp() {
    HomeActivity(
        listOf(
            ViewModel.Home.Project
                .newBuilder()
                .setID(newUUIDasByteString())
                .setName("Test project 1")
                .build(),
            ViewModel.Home.Project
                .newBuilder()
                .setID(newUUIDasByteString())
                .setName("Test project 2")
                .build()
        )
    )
}
