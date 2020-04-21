package com.example.todolist.model

import androidx.compose.Model
import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.model.ViewModelOuterClass.Location
import com.example.todolist.model.ViewModelOuterClass.LocationType
import com.example.todolist.model.CommandOuterClass.Command
import com.example.todolist.model.CommandOuterClass.CommandType
import com.example.todolist.services.GoBinding
import com.example.todolist.data.newUUIDasByteString
import com.google.protobuf.ByteString

interface HomeModelBase {
    var data: ViewModel.Home
    fun load()
    fun addProject(projectName: String)
    fun removeProject(projectId: ByteString)
}

@Model
class HomeModel : HomeModelBase {
    override var data: ViewModel.Home = ViewModel.Home.getDefaultInstance()

    override fun load() {
        val data = GoBinding.read(
            type = LocationType.Home,
            payload = Location.Home.newBuilder().build()
        )

        this.data = ViewModel.Home.parseFrom(data)
    }

    override fun addProject(projectName: String) {
        val newProject = Command.NewProject
            .newBuilder()
            .setName(projectName)
            .setProjectID(newUUIDasByteString())
            .build()

        GoBinding.execute(CommandType.NewProject, newProject)
        load()
    }

    override fun removeProject(projectId: ByteString) {
        val removableProject = Command.DeleteProject
            .newBuilder()
            .setProjectID(projectId)
            .build()

        GoBinding.execute(CommandType.DeleteProject, removableProject)
        load()
    }

}
