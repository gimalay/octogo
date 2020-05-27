package com.example.todolist.ui.home

import androidx.compose.*
import androidx.ui.core.Text
import androidx.ui.core.dp
import androidx.ui.foundation.Clickable
import androidx.ui.layout.*
import androidx.ui.text.TextStyle
import androidx.ui.text.font.FontStyle
import com.example.todolist.model.HomeFilter
import com.example.todolist.model.HomeSorter
import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.service.Navigator
import com.google.protobuf.ByteString

@Composable
fun Home(
    projects: List<ViewModel.Home.Project>,
    onAddProject: (String) -> Unit = {},
    onRemoveProject: (ByteString) -> Unit = {},
    onNavigateTo: (Navigator.Screen) -> Unit = {},
    onApplyFilter: (HomeFilter) -> Unit,
    onApplySorter: (HomeSorter) -> Unit
) {

    Column {
        HomeHeader(
            onAddProject = onAddProject,
            onApplyFilter = onApplyFilter,
            onApplySorter = onApplySorter
        )
        if (projects.isEmpty()) {
            Padding(padding = 16.dp) {
                Text(
                    text = "Projects not found",
                    style = TextStyle(fontStyle = FontStyle.Italic)
                )
            }

            return@Column
        }

        HomeBody {
            projects.forEach {
                Clickable(
                    onClick = {
                        onNavigateTo(Navigator.Screen.Article(it))
                    }
                ) {
                    Article(
                        text = it.name,
                        onCopy = { onAddProject(it.name) },
                        onRemove = { onRemoveProject(it.id) }
                    )
                }
            }
        }
    }
}
