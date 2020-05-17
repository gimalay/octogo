package com.example.todolist.ui.home

import androidx.compose.Composable
import androidx.compose.state
import androidx.compose.unaryPlus
import androidx.ui.core.EditorModel
import androidx.ui.core.dp
import androidx.ui.layout.*
import com.example.todolist.model.HomeFilter
import com.example.todolist.ui.common.ActionsRow
import com.example.todolist.ui.common.DialogButton
import com.example.todolist.ui.common.TextFieldWithHint

@Composable
fun HomeHeader(
    onAddProject: (projectName: String) -> Unit,
    onApplyFilter: (HomeFilter) -> Unit
) {
    Column {
        ActionsRow {
            FilterByProjectsAction(
                onApplyFilter = onApplyFilter
            )
            WidthSpacer(width = 12.dp)
            NewProjectAction(
                onAddProject = onAddProject
            )
        }
    }
}

@Composable
fun FilterByProjectsAction(
    onApplyFilter: (HomeFilter) -> Unit
) {
    val projectNameEditorState = +state { EditorModel() }
    val projectIdEditorState = +state { EditorModel() }

    DialogButton(
        text = "Filter",
        dialogTitle = "Filter by Project fields",
        onApply = {
            onApplyFilter(
                HomeFilter(
                    projectId = projectIdEditorState.value.text,
                    projectName = projectNameEditorState.value.text
                )
            )
        }
    ) {
        Padding(padding = 24.dp) {
            TextFieldWithHint(
                label = "ID:",
                hint = "Please, enter search keyword",
                value = projectIdEditorState.value,
                onChange = { projectIdEditorState.value = it }
            )
        }
        Padding(padding = 24.dp) {
            TextFieldWithHint(
                label = "Name:",
                hint = "Please, enter search keyword",
                value = projectNameEditorState.value,
                onChange = { projectNameEditorState.value = it }
            )
        }
    }
}

@Composable
fun NewProjectAction(onAddProject: (String) -> Unit) {
    val projectNameEditorState = +state { EditorModel() }

    DialogButton(
        text = "Create",
        disabled = projectNameEditorState.value.text.isEmpty(),
        dialogTitle = "Create Project",
        onApply = {
            onAddProject(projectNameEditorState.value.text)
            projectNameEditorState.value = EditorModel("")
        }
    ) {
        Padding(padding = 24.dp) {
            TextFieldWithHint(
                label = "Name:",
                value = projectNameEditorState.value,
                hint = "Please, enter name",
                onChange = { projectNameEditorState.value = it }
            )
        }
    }
}
