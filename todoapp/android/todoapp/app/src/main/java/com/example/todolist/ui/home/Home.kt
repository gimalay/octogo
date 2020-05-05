package com.example.todolist.ui.home

import androidx.compose.*
import androidx.ui.layout.*
import androidx.ui.material.*
import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.ui.lightThemeColors
import com.example.todolist.ui.themeTypography
import com.google.protobuf.ByteString

@Composable
fun Home(
    projects: List<ViewModel.Home.Project>,
    onAddProject: (name: String) -> Unit = {},
    onRemoveProject: (id: ByteString) -> Unit = {}
) {
    MaterialTheme(
        colors = lightThemeColors,
        typography = themeTypography
    ) {
        Column {
            HomeHeader(onAddProject)
            HomeBody {
                projects.forEach {
                    Activity(
                        text = it.name,
                        onCopy = { onAddProject(it.name) },
                        onRemove = { onRemoveProject(it.id) }
                    )
                }
            }
        }
    }
}
