
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Server.Util;
namespace Server.Controllers
{
    [Route("api/videos")]
    public class VideosController : Controller
    {
        private VideoService _service;

        public VideosController(VideoService service)
        {
            _service = service;
        }
        [HttpGet]
        public async Task<IActionResult> GetAll()
        {
            return Ok(await _service.GetAll());
        }

        [HttpGet("{id}")]
        public async Task<IActionResult> Get(int id)
        {
            return Ok(await _service.GetById(id));
        }

        [HttpDelete("{id}")]
        public async Task<IActionResult> Delete(int id)
        {
            if (await _service.DeleteById(id)) return Ok();
            else return NotFound();
        }
    }
}
