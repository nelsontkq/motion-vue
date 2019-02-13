
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Server.Util;
namespace Server.Controllers
{
    public class VideosController : Controller
    {
        public VideosController(VideoRetriever retriever)
        {

        }
        [HttpGet]
        public async Task<IActionResult> GetAll()
        {
            return Ok();
        }

        [HttpGet("{id}")]
        public async Task<IActionResult> Get(int id)
        {
            return Ok();
        }
    }
}
